package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"crypto/rand"
	"encoding/hex"
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

func main() {
	l6logins = []l6Login{
		l6Login{"ghoning", "prog50%", getSessionId(), "Dit profiel heeft nog geen data.", "prog50%, gerard123"},
		l6Login{"tmsbrg", "password123", getSessionId(), "Maker van de hacking challenge", "password123, koelkast88"},
		l6Login{"jb", "jb", getSessionId(), "Just in", "jb, nietradenalstublieft"},
		l6Login{"admin", "admin", getSessionId(), "Administrator voor Level 6", "admin, DHXoKytdyOlDSs2YXft6dynImvDfhdBm"},
	}
	http.HandleFunc("/level4/MHjgkjTp0APwXa50BnqItTJ36EHdLjkU/login.go", l4_login_thompson)
	http.HandleFunc("/level5/aWV2NNllDyktjkwlIfvu53jIBmIocRSb/viewprofile.go", l5_viewprofile)
	http.HandleFunc("/level6/vvJBa5GqajeBGhyUczJVPIYyOtlWtTd3/inloggen.go", l6_login_wa)
	http.HandleFunc("/level6/vvJBa5GqajeBGhyUczJVPIYyOtlWtTd3/wa.go", l6_wa)
	log.Println("Starting backend...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
