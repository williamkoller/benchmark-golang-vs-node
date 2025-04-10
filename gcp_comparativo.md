
# 💵 Comparativo de Custo: Go vs Node.js no Google Cloud (GCP)

## 🧮 Visão Geral

| **Cenário** | **Node.js** | **Go (Golang)** |
|-------------|-------------|-----------------|
| **Cold Start (Cloud Functions)** | 150–400ms | 10–60ms ⚡ |
| **Consumo de Memória** | 100–200MB | 10–50MB |
| **Uso de CPU (Cloud Run)** | Alto | Baixo a moderado |
| **Tamanho do Deploy** | 2–10MB+ (zip) | 2–3MB (binário) |
| **Startup (Cloud Run)** | 1–3s | 0.3–1.5s |
| **Imagem Docker típica** | 150–400MB | 10–30MB |
| **Custo Cloud Functions (1M execs, 128MB/200ms)** | ~$2.30 USD | ~$0.60 USD |
| **Custo Cloud Run (0.25vCPU, 512MB RAM)** | ~$11–13 USD | ~$8–10 USD |

---

## 🚀 Cloud Functions

### Node.js
- Cold start entre 150ms e 400ms
- Consumo mais alto de RAM
- Recomendado para desenvolvimento rápido com JS

### Go
- Cold start rápido
- Ótimo para workloads de baixa latência
- Menor custo operacional

---

## 🚢 Cloud Run

### Node.js
- Boa integração e suporte nativo
- Alto uso de memória com workloads intensos

### Go
- Binários eficientes
- Excelente escalabilidade com baixo custo

---

## 💵 Simulação de Custos

**1 milhão de execuções/mês, 128MB RAM, 200ms**

- Node.js: **~$2.30/mês**
- Go: **~$0.60/mês**

---

## ✅ Quando usar

| Situação | Melhor escolha |
|----------|----------------|
| Desenvolvimento rápido | Node.js |
| Performance / custo | Go |
| Escalabilidade serverless | Go |
| MVP com time JS | Node.js |
| Containers leves e performáticos | Go |

---

Feito com foco em custo, escalabilidade e performance ☁️🚀
