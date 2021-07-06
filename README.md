# protodoc-serv

## How to build and deploy

1. Build the portal. The portal has static web pages and some handlers.

```
cd portal
yarn build
```

2. Generate proto files and static assets bundle.

```
cd ..
go generate -x ./...
```

3. Build the project.

```
go build
```

4. Deploy and run.

```
./protodoc-serv
```
