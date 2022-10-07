Clean Architecture adalah arsitektur perangkat lunak dan Domain Driven Design adalah teknik desain perangkat lunak.
 
Kendala sebelum merancang clean architecture yaitu : 
1. Independent of Frameworks  
2. Testable 
3. Independent of UI 
4. Independent of Database 
5. Independent of any External 
Keuntungannya yaitu : 
1. Struktur standar, sehingga mudah untuk menemukan jalan dalam project 
2. Perkembangan yang lebih cepat dalam jangka panjang 
3. Mocking dependencies menjadi sepele dalam tes unit 
4. Mudah beralih dari prototipe ke solusi yang tepat (misalnya mengubah penyimpanan dalam memori ke database SQL). 
 
CA Layer : 
- Entities Layer (Optional) 
- Use Case - Domain Layer 
- Controller - Presentation Layer 
- Drivers - Data Layer

Use Case = Layer ini merupakan layer yang akan bertugas sebagai pengontrol, yakni menangangi business logic pada setiap domain. Layer ini juga bertugas memilih repository apa yang akan digunakan, dan domain ini bisa memiliki lebih dari satu repository layer.