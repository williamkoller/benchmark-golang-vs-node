
# 💵 Comparativo de Custo: Go vs Node.js na AWS

## 🧮 Visão Geral

| **Cenário** | **Node.js** | **Go (Golang)** |
|-------------|-------------|-----------------|
| **Cold Start (Lambda)** | 100–300ms | 5–50ms ⚡ |
| **Consumo de Memória (Lambda)** | 100–200MB | 10–50MB |
| **Uso de CPU** | 60–100% | 30–60% |
| **Tamanho do Deploy (Zip)** | 2–10MB+ | 1–2MB |
| **Startup em ECS/Fargate** | 2–5s | 0.5–2s |
| **Escalabilidade Horizontal** | Boa, mas com overhead | Excelente |
| **Imagem Docker típica** | 150–400MB | 10–30MB |
| **Custo Lambda (1M execs, 128MB/200ms)** | ~$2.10 USD | ~$0.50 USD |
| **Custo Fargate (0.25vCPU, 512MB)** | ~$10–12 USD | ~$8–9 USD |

---

## 🚀 Cenário 1: AWS Lambda

### Node.js
- Cold start entre 100ms e 300ms
- Consumo maior de memória
- Ideal para APIs leves e projetos JS

### Go
- Cold start muito rápido
- Executa com baixa latência
- Custo por execução até **3–4x menor**

---

## 🚢 Cenário 2: ECS / Fargate

### Node.js
- Imagens leves com Alpine
- Depende de runtime Node.js

### Go
- Binário único
- Imagens Docker de 10–30MB com distroless ou scratch

---

## 💵 Simulação de Custos com Lambda

**1 milhão de execuções/mês, 128MB RAM, 200ms**

- Node.js: **~$2.10/mês**
- Go: **~$0.50/mês**

---

## ✅ Quando usar

| Situação | Melhor escolha |
|----------|----------------|
| API simples com JS | Node.js |
| Alta performance / baixo custo | Go |
| Apps CPU/memória intensiva | Go |
| MVP / time JS | Node.js |
| Lambda serverless em escala | Go |

---

Feito para quem se importa com **performance e custo** 💡
