# assumes installed:
# nginx letsencrypt python3-certbot-nginx
# assumes letsencrypt configured
# assumes nginx configured (see config files)

echo 'Adding user...'
LC_ALL=C adduser --system --no-create-home --gecos '' --shell /bin/false --disabled-login --group weekend

echo 'Deploying site...'
mkdir -p /var/www/weekendschool
cd /var/www/weekendschool
rm -rf ./*
tar xf /home/ubuntu/weekend.tar.gz
chown -R root: .

echo 'Installing backend...'
install -m 664 weekend.service /etc/systemd/system
systemctl daemon-reload

echo 'Running backend & reloading nginx...'
systemctl enable weekend
systemctl restart weekend
systemctl reload nginx
