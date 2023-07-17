const xhr = new XMLHttpRequest();

xhr.open("GET", "your-url", true);
// param 1 method
// param 2 your url / link sumber(data by url)
// param 3 true asynchronus / false synchronus

// "GET (ngambil data)", "POST (ngirim data)", "PATCH (update data)", "DELETE (untuk mendelete data)"

xhr.onload = function () {}; // ---> untuk mengecek status;
xhr.onerror = function () {}; // ---> untuk menanpilkan error ketika request baik internet mati atau server error
xhr.send(); // ---> untuk mengirim configur yang kita bikin
