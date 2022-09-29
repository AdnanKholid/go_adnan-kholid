Middleware adalah entitas yang terhubung ke server request ataupun response server. Middleware mengizinkan user untuk mengimplementasikan berbagai fungsi di sekitar permintaan masuk http ke server dan respons keluar. Dengan middleware, kita dapat memfilter user dan memproteksi sebuah halaman agar hanya bisa diakses oleh user yang mempuyai peran tertentu.
Auth Middleware adalah alah satu middleware yang datang automatis dengan pada autentikasi di laravel. Middleware dijadikan jembatan antara user dan request urlnya, kita bisa mengecek apakah user ini mendapat izin atau tidak
JWT Middleware adalah salah satu standar JSON (RFC 7519) untuk keperluan akses token. Token dibentuk dari kombinasi beberapa informasi yang di-encode dan di-enkripsi. Informasi yang dimaksud adalah header, payload, dan signature.

-- Setup Middleware Echo 
Echo#Pre(),                          
- HTTPSRedirect 
- HTTPSWWWRedirect 
- WWWRedirect 
- AddTrailingSlash 
- RemoveTrailingSlash 
- MethodOverride 
- Rewrite 
 
Echo#Use(),
- BodyLimit 
- Logger 
- Gzip 
- Recover 
- BasicAuth 
- JWTAuth 
- Secure 
- CORS 
- Static