
# 🐳☸️ Comparativo: Docker + Kubernetes – AWS vs GCP

Este comparativo avalia desempenho, custo e usabilidade de **containers Docker** orquestrados com **Kubernetes** em **AWS (EKS)** e **Google Cloud (GKE)**.

---

## ⚙️ Visão Geral

| Fator                         | AWS (EKS)                         | Google Cloud (GKE)                |
|------------------------------|-----------------------------------|-----------------------------------|
| **Facilidade de Setup**      | Médio (mais manual)               | Muito fácil (wizard GUI + CLI)   |
| **Integração com Docker Hub**| Sim                               | Sim                               |
| **Gerenciamento via UI**     | Limitado (via EKS Console)        | Completo (via GCP Console)        |
| **Auto Scaling**             | Manual com configurações extras   | Simples com poucos cliques       |
| **Cluster Autoscaler**       | Sim, mas exige setup adicional    | Ativado com 1 flag                |
| **Preço por nó (n1-standard-2)** | ~$27/mês + EKS control plane ($74) | ~$27/mês (sem custo extra do GKE Autopilot) |
| **EKS Control Plane**        | ~$74/mês                          | $0 (Autopilot) / $0.10 por pod/h (Standard) |
| **Logging e Monitoramento**  | CloudWatch (config extra)         | Stackdriver (nativo)              |
| **Performance Geral (latência)** | Muito boa                        | Excelente (menor cold start)      |
| **Deploys com Helm/ArgoCD**  | Sim                               | Sim                               |

---

## 💰 Simulação de Custos Mensais (3 nós, n1-standard-2)

| Item                    | AWS (EKS)     | GCP (GKE Autopilot) |
|-------------------------|---------------|----------------------|
| Custo por nó (3x)       | $27 × 3 = $81 | ~$0 (gerenciado por pod) |
| Custo do control plane  | $74           | Incluso              |
| **Total Estimado**      | **~$155/mês** | **~$90–110/mês**     |

---

## ✅ Quando usar cada um?

| Situação                                     | Melhor opção       |
|---------------------------------------------|--------------------|
| Integração com infraestrutura AWS existente  | AWS (EKS)          |
| Time pequeno e setup rápido                 | GCP (GKE Autopilot)|
| Maior controle da rede/VPC                  | AWS (EKS)          |
| Menor custo inicial                         | GCP (Autopilot)    |
| Monitoramento nativo fácil                  | GCP                |
| Projetos altamente customizados             | AWS (EKS)          |

---

## 🧠 Conclusão

- **GKE Autopilot** é ideal para times que querem simplicidade, baixo custo e deploy rápido.
- **EKS** oferece mais controle, integração com serviços AWS e flexibilidade para ambientes complexos.

Feito para quem orquestra containers com propósito 🚀🐳☁️
