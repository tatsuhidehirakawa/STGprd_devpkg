## Personal Development Practice Repository (use STGPRD environment)
1. Overview  
It is a full-stack package for web development, focusing on the technology stack presented in the next chapter. It is optimized primarily as a template for creating large-scale CtoC sites. In particular, it is RESTful, does not use CSS-in-JS, minimizes Paas dependencies, emphasizes orthogonality and loose coupling, static typing, and Docker usage.  

2. Technology stacks.  
Frontend: TypeScript, React.js, Sass, Html, AtomicDesign, MindBEMding, MUI.  
Backend: Go(Gin,sqlc), PostgreSQL.  
Development environment: Docker, GNU Make, Air(Golang), VSCode, Git&Github, A5:SQL Mk-2, Windows10, macOS Monterey, Mermaid.js.  

3. Boot procedure.  
Please check the Makefile, that's all.  
(Git and Docker are required at a minimum, and use of GNU Make is recommended).  

4. (Warning) Handling of Version Control Files.  
This package does not include the version control files "package.json", "go.mod", and "go.sum". If you use this package, please prepare these files by yourself. Refer to "Makefile" for information on how to apply the version control files.  

5. Considerations  
(1)  
(2)  

6. Directory structure diagram (overall/excerpt)  
```
myportfolio_k
  ├─k_01_src                 # Source codes.
  |   ├─110wbs_src
  |   ├─124api_src
  |   └─134dbs_src
  ├─k_02_dev                 # Management files for Development.(REST)
  |   ├─110wbs_dev
  |   |   └─Dockerfile
  |   ├─124api_dev
  |   |   └─Dockerfile
  |   ├─134dbs_dev
  |   ├─docker-compose.yml
  |   └─Makefile
  ├─k_05_tst                 # Management files for Test.
  |   ├─110wbs_tst
  |   ├─124api_tst
  |   ├─134dbs_tst
  |   ├─docker-compose.yml
  |   └─Makefile
  ├─k_08_stg                 # Management files for Stage.
  |   └─Makefile
  ├─Makefile                 # Control files for Repository.
  └─README.md
```
7. Server configuration diagram (overall)  
```
REST(ULCODC$SS)_Architecture
+--------------+  +---------------------------------------------------+
|              |  |          +-----------+                Sorce codes |
|              |  |          | *110/src  |                            |
|              |  |          |           |                            |
|              |  |          |React,Sass +------+                     |
|              |  |          |TypeScript |      |                     |
|              |  |          |           /      |                     |
|              |  |          +----------/       |                     |
|              |  |          +-----------+      |       +-----------+ |
|              |  |          | *124/src  |      |       | *134/src  | |
|              |  |          |           |      |       |           | |
|              |  |          |  Golang   +---+  |  +----+ PostgreSQL| |
|              |  |          |(Gin/sqlc) |   |  |  |    |           | |
|              |  |          |           /   |  |  |    |           / |
|              |  |          +----------/    |  |  |    +----------/  |
|              |  +--------------------------|--|--|------------------+
|              |  |          +-----------+   |Development environment |
|              |  |          | *110/dev  |   |  |  |                  |
|              |  |          | WebServer |   |  |  |                  |
|              |  |          |React,Sass +<--|--+  |                  |
|              |  |          |TypeScript |   |  |  |                  |
|              |  | +---#3000|#3000      |   |  M:MultiStageBuild     |
|  For develop |  | |        +-----------+   |  |  |                  |
|localhost:3000|<---+        +-----------+   |  |  |    +-----------+ |
|"make init.d" |  | |        | *124/dev  |   |  |  |    | *134/dev  | |
|              |  | |        | APIServer |   |  |  |    |  DBServer | |
|              |  | |        |  Golang   +<--+  |  +--->+ PostgreSQL| |
|              |  | |        |   (Air)   |   |  |  |    |           | |
|              |  | +-->#8080|#8080 #8080|#3000<-->#5432|#5432      | |
|              |  |          +-----------+   |  |  |    +-----------+ |
|    CLIENT    |  +--------------------------M--|--|------------------+
|   (Browser)  |  |          +-----------+   |  |  | Test environment |
|              |  |          | *110/tst  |   |  |  |                  |
|              |  |          | WebServer |   |  |  |                  |
|              |  |          | Html, CSS +<-----+  |                  |
|              |  |          |           |   |     |                  |
|              |  | +---#3001|#3000      |   |     D:DumpAndRestore   |
|   For test   |  | |        +-----------+   |     |                  |
|localhost:3001|<---+        +-----------+   |     |    +-----------+ |
|"make init.t" |  | |        | *120/tst  |   |     |    | *130/tst  | |
|              |  | |        | APIServer |   |     |    |  DBServer | |
|              |  | |        |  Golang   +<--+     +--->+ PostgreSQL| |
|              |  | |        | (Binary)  |              |           | |
|              |  | +-->#8081|#8080 #8080|#3001<-->#5433|#5432      | |
|              |  |          +-----------+              +-----------+ |
+--------------+  +---------------------------------------------------+
```


<!--
使いやすさを優先せず、１０年後でも理解できるコードを！

-->
