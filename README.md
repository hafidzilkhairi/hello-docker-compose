# hello-docker-compose

<Table>
<thead>
<tr>
<td>
Nama
</td>
<td>
NIM
</td>
</tr>
<thead>

<tr>
<td>
Hafidzil Khairi
</td>
<td>
1301160171
</td>
</tr>

<tr>
<td>
Yola Adipratama
</td>
<td>
1301144156
</td>
</tr>
</Table>

Repository ini berisi file-file berupa web application dengan backeng menggunakan golang

Web application tersebut dibungkus kedalam container menggunakan docker dan kubernetes

## Requirement
- Docker
- kubectl, minikube

## Installation
- Clone repository ini
- Ubah working directory ke directory repository yang telah anda clone
- jalankan ```./setup```
- Jika anda sudah menjalankan setup, untuk menjalankan kedua dan seterusnya hanya jalankan perintah ```docker-compose up```

## Hasil
1. Docker
 - Jalan kan ```./setup``` terlebih dahulu

<img src="https://github.com/hafidzilkhairi/hello-docker-compose/blob/master/image/Docker.png?raw=true"/>

- Selanjutnya container dapat dihenti/jalankan menggunakan perintah ```docker-compose (up/down)```
<img src="https://github.com/hafidzilkhairi/hello-docker-compose/blob/master/image/Docker-compose.png?raw=true" />

2. Kubernetes
- file docker-compose.yml dapat diconvert ke dalam file kubernetes dengan ```kompose convert -f <nama-file>``` dan menjalankan dengan minikube
<img src="https://github.com/hafidzilkhairi/hello-docker-compose/blob/master/image/Kubernetes.png?raw=true" />

3. Hasil
- Setelah anda menjalankan perintah sebelumnya, untuk melihat hasil anda dapat mengunjungi url "localhost:8000" atau "127.0.0.1:8000"
<img src="https://github.com/hafidzilkhairi/hello-docker-compose/blob/master/image/hasil.png?raw=true" />
