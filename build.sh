#! /bin/bash
case $1 in
    mac )
        # 构建web和其他services
        echo "start build mac ..."

        cd $GOPATH/src/streaming/api
        go build -o ../bin/api

        cd $GOPATH/src/streaming/scheduler
        go build -o ../bin/scheduler

        cd $GOPATH/src/streaming/streamserver
        go build -o ../bin/streamserver

        cd $GOPATH/src/streaming/web
        go build -o ../bin/web

        cp $GOPATH/src/streaming/config/conf.json     $GOPATH/src/streaming/bin/
    ;;

    linux  )
        # 构建web和其他services
         echo "start build linux ..."

        cd $GOPATH/src/streaming/api
        env GOOS=linux GOARCH=amd64 go build -o ../bin/api

        cd $GOPATH/src/streaming/scheduler
        env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

        cd $GOPATH/src/streaming/streamserver
        env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver

        cd $GOPATH/src/streaming/web
        env GOOS=linux GOARCH=amd64 go build -o ../bin/web

        cp $GOPATH/src/streaming/config/conf.json     $GOPATH/src/streaming/bin/
    ;;
    *)

    echo "usage: build[linux|mac]"
esac