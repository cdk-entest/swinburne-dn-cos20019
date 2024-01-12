import os

# parameters
REGION = "us-east-1"
ACCOUNT = os.environ["ACCOUNT_ID"] 

# delete all docker images
os.system("sudo docker system prune -a")

# build go-app image
os.system("sudo docker build -t go-app . ")

#  aws ecr login
os.system(f"aws ecr get-login-password --region {REGION} | sudo docker login --username AWS --password-stdin {ACCOUNT}.dkr.ecr.{REGION}.amazonaws.com")

# get image id
IMAGE_ID=os.popen("sudo docker images -q go-app:latest").read()

# tag go-app image
os.system(f"sudo docker tag {IMAGE_ID.strip()} {ACCOUNT}.dkr.ecr.{REGION}.amazonaws.com/go-app:latest")

# create ecr repository
os.system(f"aws ecr create-repository --registry-id {ACCOUNT} --repository-name go-app")

# push image to ecr
os.system(f"sudo docker push {ACCOUNT}.dkr.ecr.{REGION}.amazonaws.com/go-app:latest")

# run locally to test
# os.system(f"sudo docker run -d -p 3000:3000 go-app:latest")