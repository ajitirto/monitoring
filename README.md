## 🚀 Memulai (Getting Started)

### Prasyarat

Pastikan telah menginstal tool berikut:
1.  **Docker**
2.  **Docker Compose**
2.  **Make**

### 1. Menjalankan prometheus & grafana di dalam docker

```bash
make install-prometheus
```

### 2. Menjalankan applikasi

```bash
make run
```

![app-go](docs/app-go.png)


### 3. Melihat /metrics di applikasi


![metric](docs/metrics.png)

![myapp](docs/myapp.png)


### 4. result aplikasi  

![result](docs/grafana.png)