
# ⚙️ Executando Node.js e Go com Docker + Kubernetes localmente

Este guia mostra como rodar aplicações **Node.js** e **Go** usando **Docker** com **Kubernetes local (minikube ou kind)**.

---

## 🧱 Estrutura do Projeto

```
project/
├── go-app/
│   ├── main.go
│   ├── Dockerfile
│   └── k8s.yaml
├── node-app/
│   ├── index.js
│   ├── Dockerfile
│   └── k8s.yaml
└── README.md
```

---

## 🐳 Dockerfile - Go

```Dockerfile
# go-app/Dockerfile
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o server .

FROM scratch
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
```

---

## 🐳 Dockerfile - Node.js

```Dockerfile
# node-app/Dockerfile
FROM node:20-alpine
WORKDIR /app
COPY . .
RUN npm install
CMD ["node", "index.js"]
```

---

## ☸️ Kubernetes Manifest (k8s.yaml)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
        - name: app
          image: <your-image-name>
          ports:
            - containerPort: 3000
---
apiVersion: v1
kind: Service
metadata:
  name: app-service
spec:
  selector:
    app: myapp
  ports:
    - protocol: TCP
      port: 80
      targetPort: 3000
  type: NodePort
```

---

## 🚀 Rodando localmente

### 1. Build das imagens

```bash
# Go
cd go-app
docker build -t go-app .

# Node.js
cd ../node-app
docker build -t node-app .
```

### 2. Rodando com Minikube

```bash
minikube start
eval $(minikube docker-env)

# deploy apps
kubectl apply -f go-app/k8s.yaml
kubectl apply -f node-app/k8s.yaml

# acessar
minikube service app-service
```

---

## ✅ Conclusão

Agora você tem ambos os serviços **containerizados e orquestrados via Kubernetes localmente**, ótimo para benchmarking, testes e comparações realistas antes de subir na nuvem.

Feito com containers e clusters na raiz 🐳☸️🧪
