#.PHONY:

#---[ 1. system boot ]-------------------------------------------------------

init.all: ## Initialize and booting all containers and services.
	make init.dev
	make init.tst
	# make init.stg

init.dev: ## Booting only frontend devlopment containers.
	cd s_03_dev && make init.dev

init.tst: ## Booting only backend testing containers.
	cd s_05_tst && make init.tst

init.stg: ## Booting only backend staging containers.
	cd s_06_tst && make init.stg

#---[ 3. Tst env boot ]------------------------------------------------

boot.all:
	make boot.dev
	make boot.tst
	# make boot.stg

boot.dev:
	cd s_03_dev && make boot.dev

boot.api.a:
	docker build -f s_05_tst/120api_tst/Dockerfile -t 120api_tst s_01_src/124api_src/src
	cd s_05_tst && docker compose up 120api_tst

boot.api.b:
	cp -r s_01_src/124api_src/src s_05_tst/120api_tst
	cd s_05_tst && docker compose up 120api_tst

#---[ 3. DB backup and restore ]------------------------------------------------

pg_dump: ## Dump db backupfile to local.(cf."meke pg_dump container_name=134dbs_dev")
	@docker exec -it "$(container_name)" /usr/src/dbstore/pg_dump.sh "$(container_name)" bash

pg_restore: ## Restore by local file.(cf."make pg_restore container_name=130dbs_tst file_name=sample.sql")
	@docker exec -it "$(container_name)" /usr/src/dbstore/pg_restore.sh "$(file_name)" bash

#---[ 4. Hot-reload ]-------------------------------------------------------------


#---[ 5. Validation ]----------------------------------------------------------------

look.crlf.wbs:
	@cd s_03_dev/110wbs_dev && cat -e *.sh

look.crlf.api:
	@cd s_03_dev/124api_dev && cat -e *.sh

rplccrlf/all:
	@cd s_03_dev/110wbs_dev && sed -i 's/\r//' *.sh
	@cd s_03_dev/124api_dev && sed -i 's/\r//' *.sh

yyyyyyy:
	@docker exec -it k_904set_dev-214api_dev-1 bash

clonenv:
	@git clone https://github.com/tatsuhidehirakawa/STGprd_devpkg.git

persist:
	@cd s_03_dev && 110wbs_dev && rm strset.sh
	@cd ../../s_01_src/k110wbs && git add package.json, package-lock.json

#---[ 8. Git add ]-------------------------------------------------------------

git.add.norm:  ## Addition of git (cf. make git.add.norm)
	@git status
	@git add Makefile
	@git status

git.add.param:  ## Addition of git (cf. make git.add.param)
	@git status
	@git add s_01_src/110wbs_src/package.json
	@git add s_01_src/110wbs_src/package-lock.json
	@git add s_01_src/124api_src/src/mod.go
	@git add s_01_src/124api_src/src/sqlc/build_sqlc/db.go
	@git add s_01_src/124api_src/src/sqlc/build_sqlc/go.mod
	@git add s_01_src/124api_src/src/sqlc/build_sqlc/models.go
	@git add s_01_src/124api_src/src/sqlc/build_sqlc/query.sql.go
	@git status

#---[ 9. memo ]----------------------------------------------------------------

#---[ 9.1. docker ]------------------------------------------------------------
# https://zenn.dev/sickleaf/articles/99884a12b0489cf21d45
	@docker ps -a | grep dbs | awk '$0=$NF'


#---[ 10. make test ]----------------------------------------------------------

vall:
	@make -v && git -v && docker -v

test.s:
	@touch test.txt

test.e:
	@rm test.txt


exec:
	@docker exec -it 134dbs_dev bash



debug: ## Debug code removal




init.dev.bake: ## Under construction.
	@docker buildx bake --file docker-bake.hcl STGprd_devpkg_dev STGprd_devpkg_tst



chmod:
	@cd s_03_dev && cd 124api_dev && ls -la && chmod 775 pkgset.sh && ls -la
	@cd s_03_dev && cd 110wbs_dev && ls -la && chmod 775 strset.sh && ls -la
	@cd s_05_tst && cd 130dbs_tst && ls -la && chmod 775 pg_dump.sh && ls -la
	@cd s_05_tst && cd 130dbs_tst && ls -la && chmod 775 pg_restore.sh && ls -la
