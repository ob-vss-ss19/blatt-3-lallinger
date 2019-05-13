## Ausführen mit den Docker Images des Jenkins-Servers

-   ein (Docker)-Netzwerk `actors` erzeugen

    ```
    docker network create actors
    ```

-   Starten des Tree-Services und binden an den Port 8090 des Containers mit dem DNS-Namen
    `treeservice` (entspricht dem Argument von `--name`) im Netzwerk `actors`:

    ```
    sudo docker run --rm --net actors --name treeservice terraform.cs.hm.edu:5043/ob-vss-ss19-blatt-3-lallinger:develop-treeservice --bind="treeservice.actors:8090"
    ```

-   Starten des Tree-CLI, Binden an `treecli.actors:8091` und nutzen des Services unter
    dem Namen und Port `treeservice.actors:8090`:

    ```
    sudo docker run --rm --net actors --name treecli terraform.cs.hm.edu:5043/ob-vss-ss19-blatt-3-lallinger:develop-treecli --bind="treecli.actors:8091" --remote="treeservice.actors:8090" trees
    ```
-   Hier sind wieder die beiden Flags `--bind` und `--remote` beliebig gewählt und
    in der Datei `treeservice/main.go` implementiert. Der Befehl "trees" zeigt eine Liste aller verfügbaren IDs. Weiter unten können sind alle verfügbaren Befehle und deren Gebrauch zu finden.
    durch aufrufen der treecli mit dem flag `--help` wird die gleiche Befehlsliste ausgegeben.     

-   Zum Beenden, killen Sie einfach den Tree-Service-Container mit `Ctrl-C` und löschen
    Sie das Netzwerk mit

    ```
    docker network rm actors
    ```


-   Befehlsübersicht für die treecli:
    ```
    treecli [FLAGS] COMMAND [KEY/SIZE] [VALUE]
    FLAGS
      -bind string
            Bind to address (default "localhost:8092")
      -id int
            tree id (default -1)
      -no-preserve-tree
            force deletion of tree
      -remote string
            remote host:port (default "127.0.0.1:8093")
      -token string
            tree token
    
    COMMAND
      newtree SIZE
            Creates new tree. SIZE parameter specifies leaf size (minimum 1). Returns id and token
      insert KEY VALUE
            Insert an integer KEY with given string VALUE into the tree. id and token flag must be specified
      search KEY
            Search the tree for KEY. Returns corresponding value if found. id and token flag must be specified
      remove KEY
            Removes the KEY from the tree. id and token flag must be specified
      traverse
            gets all keys and values in the tree sorted by keys. id and token flag must be specified
      trees
            Gets a list of all available tree ids
      delete
            Deletes the tree. id and token flag must be specified, also no-preserve-tree flag must be set to true
    
    Example:
      treecli newtree 2
      treecli --id=0 --token=d57a23df insert 42 "the answer"
    ```

## Compilieren und Ausführen der Sourcen

-   Herunterladen des Repositorys:
    ```
    git pull https://github.com/ob-vss-ss19/blatt-3-lallinger
    ```
    
-   Compilieren und Starten des treeservice:
    ```
    cd treeservice
    go build
    treeservice.exe
    ```    
-   Standardmäßig wird zum binden `localhost: 8093` verwendet, dies kann über `--bind` geändert werden.

-   Compilieren und Starten der treecli:
    ```
    cd treecli
    go build
    treecli.exe trees
    ```    
-   Standardmäßig wird zum binden `localhost: 8092` verwendet, dies kann über `--bind` geändert werden.

-   Die Befehlsübersicht ist weiter oben zu finden oder kann über das flag `--help` angezeigt werden.
