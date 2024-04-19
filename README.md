# Sistem Bilgisi API

Bu basit Go programı, sistem hakkında temel bilgileri sağlayan bir API sunar.

## Kullanım

- Program, `/systeminfo` endpoint'ına yapılan GET isteklerini dinler.
- Endpoint'a yapılan istekler, sistem bilgilerini JSON formatında döndürür.
- Program, sunucunun hostname'ini, işletim sistemini ve mimarisini döndürür.
- Program, 8080 portunda bir HTTP sunucusu olarak çalışır.

## Dikkat

- Program, sunucunun işletim sistemi ve mimarisini çevresel değişkenlerden alır.
- Programı çalıştırmadan önce, gerekli izinleri ve bağımlılıkları sağladığınızdan emin olun.

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Daha fazla bilgi için `LICENSE` dosyasını inceleyebilirsiniz.
