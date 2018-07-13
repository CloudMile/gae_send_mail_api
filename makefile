project_id = $(PROJECT_ID)
from_mail = $(FROM)
version = $(shell date +"%y%m%d%H%M%S")

help:
	@echo 'Support Sub Commands:'
	@echo ''
	@echo 'Deploy, you need to setup PROJECT_ID and FROM'
	@echo ''
	@echo '    $$ make deploy'
	@echo ''

deploy:
	@echo ''
	@echo 'Start to deploy'
	@echo ''

	sed -i '' 's/<YOUR_GAE_MAIL_SENDER>/$(from_mail)/' ./main/app.yaml
	gcloud app deploy -q --no-stop-previous-version --no-promote --project=$(project_id) --version=$(version) ./main/app.yaml
	sed -i '' 's/$(from_mail)/<YOUR_GAE_MAIL_SENDER>/' ./main/app.yaml
