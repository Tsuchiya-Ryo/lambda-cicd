```sh
cd infra/
terraform init
terraform plan -var-file ${.tfvars-file}
terraform apply -var-file ${.tfvars-file}

# ecrにpushされているimageを予め消しておくこと
# terraform destroy -var-file ${.tfvars-file}
```