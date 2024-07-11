# Deplong Golang CI/CD to Google Cloud Platform

This is a simple Golang Model-Controller template using [Functions Framework for Go](https://github.com/GoogleCloudPlatform/functions-framework-go) and mongodb.com as the database host. It is compatible with Google Cloud Function CI/CD deployment.

Start here: Just [Fork this repo](https://github.com/gocroot/gcp/)

## MongoDB Preparation

The first thing to do is prepare a Mongo database using this template:

1. Sign up for mongodb.com and create one instance of Data Services of mongodb.
2. Go to Network Access menu > + ADD IP ADDRESS > ALLOW ACCESS FROM ANYWHERE  
   ![image](https://github.com/gocroot/gcp/assets/11188109/a16c5a73-ccdc-4425-8333-73c6fbf78e6d)  
3. Download [MongoDB Compass](https://www.mongodb.com/try/download/compass), connect with your mongo string URI from mongodb.com
4. Create database name iteung and collection reply  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/23ccddb7-bf42-42e2-baac-3d69f3a919f8)  
5. Import [this json](https://whatsauth.my.id/webhook/iteung.reply.json) into reply collection.  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/7a807d96-430f-4421-95fe-1c6a528ba428)  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/fd785700-7347-4f4b-b3b9-34816fc7bc53)  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/ef236b4d-f8f9-42c6-91ff-f6a7d83be4fc)  
6. Create a profile collection, and insert this JSON document with your 30-day token and WhatsApp number.  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/5b7144c3-3cdb-472b-8ab3-41fe86dad9cb)  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/829ae88a-be59-46f2-bddc-93482d0a4999)  

   ```json
   {
     "token":"v4.public.asoiduasoijfiun98erjg98egjpoikr",
     "phonenumber":"6281111222333"
   }
   ```

   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/06330754-9167-4bf4-a214-5d75dab7c60a)  

## Folder Structure

This boilerplate has several folders with different functions, such as:

* .github: GitHub Action yml configuration.
* config: all apps configuration like database, API, token.
* controller: all of the endpoints functions
* model: all of the type structs used in this app
* helper: helper folder with a list of functions only called by others file
* route: all routes URL

## Deploy Azure Function

Untuk melakukan deploy Azure Function menggunakan Golang dan GitHub Actions, Anda bisa mengikuti langkah-langkah berikut:

### Persiapan Proyek

1. **Buat Proyek Azure Function:**
   Pastikan Anda sudah mengikuti langkah-langkah di atas untuk membuat proyek Azure Function dengan Golang.

2. **Inisialisasi Git:**
   Inisialisasi repository Git di folder proyek Anda:

   ```sh
   git init
   ```

3. **Buat Repository di GitHub:**
   Buat repository baru di GitHub dan hubungkan ke repository lokal Anda:

   ```sh
   git remote add origin https://github.com/username/repository.git
   ```

### Membuat Workflow GitHub Actions

1. **Buat Folder Workflow:**
   Buat folder `.github/workflows` di dalam root proyek Anda.

2. **Buat File Workflow:**
   Buat file baru di dalam folder `.github/workflows` dengan nama `deploy.yml` dan tambahkan konfigurasi berikut:

   ```yaml
   name: Deploy Azure Function

   on:
     push:
       branches:
         - main

   jobs:
     build-and-deploy:
       runs-on: ubuntu-latest

       steps:
       - name: Checkout code
         uses: actions/checkout@v2

       - name: Set up Go
         uses: actions/setup-go@v2
         with:
           go-version: '1.16'

       - name: Install Azure CLI
         run: |
           curl -sL https://aka.ms/InstallAzureCLIDeb | sudo bash

       - name: Login to Azure
         uses: azure/login@v1
         with:
           creds: ${{ secrets.AZURE_CREDENTIALS }}

       - name: Deploy to Azure Functions
         run: |
           func azure functionapp publish <your_function_app_name>

   ```

### Menyiapkan Secrets di GitHub

1. **Buat Service Principal:**
   Buat Service Principal untuk otentikasi dengan Azure:

   ```sh
   az group list --output table
   az account list --output table

   func init --worker-runtime custom
   func new --template "HTTP trigger" --name gocroot

   az account set --subscription <your_subscription_id>

   az functionapp config appsettings set --name <your_function_app_name> --resource-group <your_resource_group> --settings FUNCTIONS_WORKER_RUNTIME=custom


   az ad sp create-for-rbac --name "myGitHubActionsServicePrincipal" --role contributor --scopes /subscriptions/<your_subscription_id>/resourceGroups/<your_resource_group> --sdk-auth
   ```

   Ini akan menghasilkan JSON output yang berisi kredensial yang diperlukan untuk login ke Azure.

2. **Tambah Secrets di GitHub:**
   Tambahkan kredensial yang dihasilkan ke secrets repository GitHub Anda:
   * Buka repository Anda di GitHub.
   * Pergi ke `Settings > Secrets > Actions`.
   * Klik `New repository secret`.
   * Nama secret adalah `AZURE_CREDENTIALS`.
   * Isi nilai secret dengan JSON output dari langkah sebelumnya.

### Menjalankan Workflow

Sekarang, setiap kali Anda push ke branch `main`, GitHub Actions akan otomatis menjalankan workflow untuk build dan deploy Azure Function Anda.

### Langkah-langkah Tambahan

* **Commit dan Push Kode:**

  ```sh
  git add .
  git commit -m "Initial commit"
  git push origin main
  ```

* **Verifikasi Deployment:**
  Setelah workflow berhasil dijalankan, Anda dapat memverifikasi deployment Anda dengan mengakses URL Azure Function yang telah Anda deploy.

Dengan mengikuti langkah-langkah di atas, Anda dapat melakukan deploy Azure Function menggunakan Golang dan GitHub Actions secara otomatis setiap kali ada perubahan di branch `main`.
