# tes_citra-digital

### 0. No 0

Buat fungsi sederhana untuk menghitung jumlah huruf hidup dan huruf mati dari suatu kalimat. Contoh:
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; input: "omama"   
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; hasil: "huruf mati: 2, huruf hidup: 2"  (huruf hidup yaitu o dan a. a tidak perlu muncul 2x)

### 1. No 1

Buat fungsi sederhana untuk mengurutkan abjad dari suatu kalimat, dengan memisahkan huruf hidup dan huruf mati. Contoh:
   <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; input: "omama"   hasil: "aaomm" ('a', dan 'o' sebagai huruf hidup dan 'm' sebagai huruf mati)
   <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; input: "osama"   hasil: "aaoms"

### 2. Webservice

* Method
    "POST"
    
* Endpoint
    "/count"
    
* Body
    { input: "string" }
    
    
### 3. Docker for no 2

* `docker build -t citra_digital_test .`
* `docker run -p 1234:1234 -t citra_digital_test`
  
  
