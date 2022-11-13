## GitHub OIDC provider
https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/configuring-openid-connect-in-amazon-web-services


ロール作成時にウェブアイデンティティを選択
- プロバイダー = token.actions.githubusercontent.com
- Audience = sts.amazonaws.com

Actionsのワークフローに以下が必要
```yml
permissions:
  id-token: write
  contents: read
```