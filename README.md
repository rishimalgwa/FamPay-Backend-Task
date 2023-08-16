# FamPay-Backend-Task

Steps to Deploy:
```
sudo docker build -t rishimalgwa/FamPay-Backend-Task .
sudo docker run -d -p 3000:3000 --name template --env-file .env rishimalgwa/FamPay-Backend-Task
```

Steps to Redeploy 
```
sudo docker build -t rishimalgwa/FamPay-Backend-Task .
sudo docker stop template
sudo docker rm template
sudo docker run -d -p 3000:3000 --name template --env-file .env rishimalgwa/FamPay-Backend-Task
```
