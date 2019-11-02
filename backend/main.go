package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"crypto/rand"
	"encoding/hex"
	"io"
	"bytes"

	"github.com/gobuffalo/packr"
)

func l4_login_thompson(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("user") != "ken" {
		http.Error(w, "Unknown user", 403)
		return
	}
	if r.FormValue("password") != "p/q2-q4!" {
		http.Error(w, "Wrong password", 403)
		return
	}
	fmt.Fprint(w, "Welcome, Ken Thompson. The password for level 5 is YVdWMk5ObGxEeWt0amt3bElmdnU1M2pJQm1Jb2NSU2I=")
}

func l5_viewprofile(w http.ResponseWriter, r *http.Request) {
	user, err := strconv.Atoi(r.FormValue("user"))
	if err != nil {
		http.Error(w, "Could not parse user number.", 403)
	} else if user == 1 {
		fmt.Fprint(w, "Username: ken\nFull name: Ken Thompson\nBirthday: February 4, 1943\nPassword: p/q2-q4!\nPersonal note: You can't trust code that you did not totally create yourself. (Especially code from companies that employ people like me.)")
	} else if user == 2 {
		fmt.Fprint(w, "Username: dmr\nFull name: Dennis Ritchie\nBirthday: September 9, 1941\nPassword: dmac\nPersonal note: C has the power of assembly language and the convenience of … assembly language. ")
	} else if user == 3 {
		fmt.Fprint(w, "Username: bwk\nFull name: Brian Kernighan\nBirthday: January 1, 1942\nPassword: /.,/.,\nPersonal note: Advice to students: Leap in and try things. If you succeed, you can have enormous influence. If you fail, you have still learned something, and your next attempt is sure to be better for it. Advice to graduates: Do something you really enjoy doing. If it isn’t fun to get up in the morning and do your job or your school program, you’re in the wrong field.")
	} else if user == 4 {
		fmt.Fprint(w, "Username: rms\nFull name: Richard Stallman\nBirthday: 1956\nPassword: <none>\nPersonal note: The use of “hacker” to mean “security breaker” is a confusion on the part of the mass media. We hackers refuse to recognize that meaning, and continue using the word to mean someone who loves to program, someone who enjoys playful cleverness, or the combination of the two.")
	} else if user == 5 {
		fmt.Fprint(w, "Username: rob\nFull name: Rob Pike\nBirthday: March 16, 1953\nPassword: vvJBa5GqajeBGhyUczJVPIYyOtlWtTd3\nPersonal note: TODO: Change the password for level 5.")
	} else {
		http.Error(w, "Unknown user", 403)
	}
}

type l6Login struct {
	User string
	Password string
	SessionId string
	UserData string
	PasswordsSaved string
}

func getSessionId() string {
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		log.Fatal("Can't read random bytes to create session IDs for Level 6. Exiting.")
	}
	return hex.EncodeToString(buf)
}

var l6logins []l6Login

func l6_login_wa(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	password := r.FormValue("password")
	for _, l := range l6logins {
		if user == l.User && password == l.Password {
			http.SetCookie(w, &http.Cookie{Name:"L6SESSID",Value:l.SessionId})
			http.Redirect(w, r, "/level6/vvJBa5GqajeBGhyUczJVPIYyOtlWtTd3/wa.go", http.StatusFound)
			return
		}
	}
	http.Error(w, "Unknown user or password.", 403)
}

func l6_wa(w http.ResponseWriter, r *http.Request) {
	var logged_in_user *l6Login = nil
	session, err := r.Cookie("L6SESSID")
	if err != nil {
		http.Redirect(w, r, "/level6/vvJBa5GqajeBGhyUczJVPIYyOtlWtTd3/inloggen.html", http.StatusFound)
		return
	}
	for _, l := range l6logins {
		if session.Value == l.SessionId {
			logged_in_user = &l
			break
		}
	}
	if logged_in_user == nil {
		http.Redirect(w, r, "/level6/vvJBa5GqajeBGhyUczJVPIYyOtlWtTd3/inloggen.html", http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Ingelogd als %s.\n\nGebruikers:\n\n", logged_in_user.User)
	for _, l := range l6logins {
		passwords := "geen toegang"
		if session.Value == l.SessionId {
			passwords = l.PasswordsSaved
		}
		fmt.Fprintf(w, "\tnaam: %s,\n\tinformatie: %s\n\topgeslagen wachtwoorden: %s\n\n", l.User, l.UserData, passwords)
	}
}

var l7_audio packr.Box

var l7_code_secret = "1442"

func l7_checkcode(code string) (num_digits_correct int) {
	num_digits_correct = 0
	if len(code) != len(l7_code_secret) {
		return
	}
	for i := 0; i < 4; i++ {
		if (l7_code_secret[i] == code[i]) {
			num_digits_correct++
		} else {
			return
		}
	}
	return
}

func l7_getpassword(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	if code == l7_code_secret {
		fmt.Fprint(w, "U3N2kWycdLiugOrRcxzkApMU1Y5LnYhs")
	} else {
		http.Error(w, "Wrong keycode", 403)
	}
}

func l7_code(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	num_digits_correct := l7_checkcode(code)

	var audiofile string
	if num_digits_correct == 0 {
		audiofile = "./0000.mp3"
	} else if num_digits_correct == 1 {
		audiofile = "./1000.mp3"
	} else if num_digits_correct == 2 {
		audiofile = "./1100.mp3"
	} else if num_digits_correct == 3 {
		audiofile = "./1110.mp3"
	} else if num_digits_correct == 4 {
		audiofile = "./1111.mp3"
	}

	audio, err := l7_audio.Find(audiofile)
	if err != nil {
		log.Println("Error: can't open audio:", err)
		http.Error(w, "Internal error", 500)
		return
	}

	w.Header().Set("Content-Type", "audio/mpeg")

	if num_digits_correct == 4 {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(403)
	}

	io.Copy(w, bytes.NewReader(audio))
}

func main() {

	http.HandleFunc("/level4/MHjgkjTp0APwXa50BnqItTJ36EHdLjkU/login.go", l4_login_thompson)

	http.HandleFunc("/level5/aWV2NNllDyktjkwlIfvu53jIBmIocRSb/viewprofile.go", l5_viewprofile)

	l6logins = []l6Login{
		l6Login{"ghoning", "prog50%", getSessionId(), "Dit profiel heeft nog geen data.", "prog50%, gerard123"},
		l6Login{"tmsbrg", "password123", getSessionId(), "Maker van de hacking challenge", "password123, koelkast88"},
		l6Login{"jb", "jb", getSessionId(), "Just in", "jb, nietradenalstublieft"},
		l6Login{"admin", "admin", getSessionId(), "Administrator voor Level 6", "admin, DHXoKytdyOlDSs2YXft6dynImvDfhdBm"},
	}
	http.HandleFunc("/level6/vvJBa5GqajeBGhyUczJVPIYyOtlWtTd3/inloggen.go", l6_login_wa)
	http.HandleFunc("/level6/vvJBa5GqajeBGhyUczJVPIYyOtlWtTd3/wa.go", l6_wa)

	l7_audio = packr.NewBox("./audio")
	http.HandleFunc("/level7/DHXoKytdyOlDSs2YXft6dynImvDfhdBm/code.go", l7_code)
	http.HandleFunc("/level7/DHXoKytdyOlDSs2YXft6dynImvDfhdBm/getpassword.go", l7_getpassword)

	log.Println("Starting backend...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
