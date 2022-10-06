## Personal Development Practice Repository (use STGPRD environment)
1. Overview  
RESTfulを意識  
CSS-in-JS非採用  
Paas依存の極小化  
直交性と疎結合、静的型付けとDocker利用を重視  

2. Technology stack  
Frontend: TypeScript、React.js、Sass、Html  
Backend: Go(Gin,sqlc)、PostgreSQL  
Development environment: Docker、VSCode、Git&Github、A5:SQL Mk-2、Windows10  

3. Boot procedure  
Please check the Makefile, that's all.

4. Branch names  
master -> Existing version.  
hotfix-100622 -> API server sorce code development.  
<!-- develop -> Next Version.  -->
<!-- release-*  -->
<!-- hotfix-100622 -->

5. Considerations  
(1)改行文字の制御に「.gitattributes」を使用  
(2)  

6. Server configuration diagram (overall)  
```
REST(ULCODC$SS)_Architecture
+--------------+  +---------------------------------------------------+
|              |  |          +-----------+    Development environment |
|              |  |          | *110_dev  |                            |
|              |  |          | WebServer |                            |
|              |  |          |React,Sass |                            |
|              |  |          |           |                            |
|              |  | +---#8080|#8080      |                            |
|              |  | |        +-----------+                            |
|              |  | |        +-----------+              +-----------+ |
|localhost:3000|<---+        | *124_dev  |              | *134_dev  | |
|              |  | |        | APIServer |              |  DBServer | |
|              |  | |        |  Golang   |              | PostgreSQL| |
|              |  | |        |(Gin/sqlc) |              |           | |
|              |  | +-->#3000|#8080 #8080|#3000<-->#5432|#5432      | |
|              |  |          +-----------+              +-----------+ |
|    CLIENT    |  +---------------------------------------------------+
|   (Browser)  |  |          +-----------+           Test environment |
|              |  |          | *110_tst  |                            |
|              |  |          | WebServer |          Ubuntu 22.04 LTS  |
|              |  |          |           |          Amazon Linux 2022 |
|              |  |          |           |                            |
|              |  | +---#8080|#8080      |                            |
|              |  | |        +-----------+                            |
|              |  | |        +-----------+              +-----------+ |
|           #80|<---+        | *120_tst  |              | *130_tst  | |
|              |  | |        | APIServer |              |  DBServer | |
|              |  | |        |  Golang   |              | PostgreSQL| |
|              |  | |        |           |              |           | |
|              |  | +-->#3000|#8080 #8080|#3000<-->#5432|#5432      | |
|              |  |          +-----------+              +-----------+ |
+--------------+  +---------------------------------------------------+
```
7. Directory structure diagram (overall/excerpt)  
```
myportfolio_k
  ├─k_01_src            # Source codes.
  ├─k_02_dev            # Management files for development.
  ├─k_03_tst            # Management files for test.
  ├─Makefile
  └─README.md
```

8. Task list  
- [x] [#2785](k_01_src/124api_src/src/main.go) : Add "post" method on "main.go".  
- [ ] [#7236](k_01_src/124api_src/src/main.go) : Add "delete" method on "main.go".  
- [ ] [#9832](k_01_src/124api_src/src/main.go) : Add "put" method on "main.go".  
- [ ] [#5982](k_01_src/124api_src/src/main.go) : Add "Environment variable" on "main.go".  
- [ ] [#8246](k_01_src/124api_src/src/main.go) : Add "Authentication mechanism" on "API".  
- [ ] [#3256](k_03_tst/120api_tst/Dockerfile) : Add test environment on "k_03_tst".(Use multistage build.)  

<!--
-->
