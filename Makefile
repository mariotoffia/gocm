generate:
	@docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`/cmselect:/work antlr/antlr4 -Dlanguage=Go cmselect.g4

clean:
	@cd cmselect;find . ! -name '*.go' ! -name '*.g4' -type f -exec rm {} +;cd -
dep:
	@echo "Removing old antlr4 docker image"
	@docker rmi $(docker images 'antlr/antlr4' -a -q)

	@echo "Cloning antlr4 docker image & building it"
	@git clone https://github.com/antlr/antlr4.git
	@cd antlr4/docker;docker build -t antlr/antlr4 --platform linux/amd64 .;cd -
	@rm -rf antlr4
	
