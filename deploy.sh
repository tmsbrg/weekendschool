echo 'Building backend...'
cd backend
go build
cd ..

echo 'Building tar...'
rm weekend.tar.gz
tar czf weekend.tar.gz *

echo 'Uploading tar...'
scp weekend.tar.gz ubuntu@hackweekend.fun:.
echo 'Uploading deployment script...'
scp deploy-remote.sh ubuntu@hackweekend.fun:deploy.sh

echo 'Running deployment script...'
echo 'Going remote...'
ssh ubuntu@hackweekend.fun 'chmod +x /home/ubuntu/deploy.sh'
ssh ubuntu@hackweekend.fun 'sudo /home/ubuntu/deploy.sh'
echo 'Finished.'
