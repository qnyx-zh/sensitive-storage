echo javajava321 | sudo -S docker stop passwd
sudo docker container rm passwd
sudo docker image rm kiririx/passwd:latest
sudo docker pull kiririx/passwd:latest
sudo docker run -p 10011:8080 -d -v ~/data/sensitive.db:/sensitive.db --name passwd kiririx/passwd:latest

#docker stop passwd
#docker container rm passwd
#docker image rm kiririx/passwd:latest
#docker pull kiririx/passwd:latest
#docker run -p 10011:8080 -d -v ~/data/sensitive.db:/sensitive.db --name passwd kiririx/passwd:latest
