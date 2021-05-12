.PHONY: aws tf clean

aws:
	GOOS=linux GOARCH=amd64 go build -o aws aws.go && zip aws.zip aws
tf:
	cd dist/terraform && terraform apply -auto-approve
clean:
	rm main.zip main
