# create network stack
aws cloudformation create-stack \
 --stack-name network-stack \
 --template-body file://1-network.yaml \
 --capabilities CAPABILITY_NAMED_IAM


aws cloudformation update-stack \
 --stack-name network-stack \
 --template-body file://1-network.yaml \
 --capabilities CAPABILITY_NAMED_IAM \
 --region us-east-1


# create web-server
aws cloudformation create-stack \
 --stack-name web-server-stack \
 --template-body file://2-web-server.yaml \
 --capabilities CAPABILITY_NAMED_IAM \
 --parameters '[{"ParameterKey":"InstanceKeyName","ParameterValue":"cos20019-2025-hai"}]' \
 --region us-east-1

# create rds mysql instance
aws cloudformation create-stack \
 --stack-name rds-instance-postgresql-stack \
 --template-body file://2-rds-instance-postgresql.yaml \
 --capabilities CAPABILITY_NAMED_IAM \
 --parameters '[{"ParameterKey":"NetworkStackName","ParameterValue":"network-database"},{"ParameterKey":"DBUsername","ParameterValue":"demo"},{"ParameterKey":"DBPassword","ParameterValue":"Password"}]'