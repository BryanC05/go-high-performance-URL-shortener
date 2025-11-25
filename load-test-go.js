import http from 'k6/http';
import { check } from 'k6';

// KONFIGURASI BEBAN:
// Kita lipat gandakan beban dari project sebelumnya.
// 100 User menembak server bersamaan selama 30 detik.
export const options = {
  vus: 100, 
  duration: '30s',
};

export default function () {
  const url = 'http://localhost:8080/shorten';
  
  // Data dummy yang dikirim
  const payload = JSON.stringify({
    url: "https://www.google.com/search?q=load+testing+golang",
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.post(url, payload, params);

  // Validasi: Harus sukses 200 OK
  check(res, {
    'status is 200': (r) => r.status === 200,
  });
}