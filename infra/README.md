```sh
cd infra/
terraform init
terraform plan -var-file ${.tfvars-file}
terraform apply -var-file ${.tfvars-file}

# ecrとs3の中身を空にしてから
# terraform destroy -var-file ${.tfvars-file}
```