// create function submitdata for send email
function submitData(jaya) {
    // menangani reload halaman
    jaya.preventDefault()

    // deklarasi variabel / tampung nilai dari inputan html ke variabel menggunakan DOM (Document Object Model)
    let name = document.getElementById("input-name").value
    let email = document.getElementById("input-email").value
    let phone = document.getElementById("input-phone").value
    let subject = document.getElementById("input-subject").value
    let message = document.getElementById("input-message").value

    // logika untuk inputan kosong
    if (name === "") {
        return alert('Name harus diisi!')
    } else if (email === "") {
        return alert('Email harus diisi!')
    } else if (phone === "") {
        return alert('Phone harus diisi!')
    } else if (subject === "") {
        return alert('Subject harus diisi!')
    } else if (message === "") {
        return alert('Message harus diisi!')
    }

    // test alert
    // alert(`name : ${name}\nemail : ${email}\nphone : ${phone}\nsubject : ${subject}\nmessage: ${message}`)

    // deklarasi emailreceiver
    const emailReceiver = "bagza@gmail.com"

    // deklarasi untuk membuat element a menggunakan DOM bertujuan untuk mengirim data form ke emailreciver
    let a = document.createElement('a')
    a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo nama saya ${name},\n${message}, silahkan kontak saya di nomor berikut : ${phone}`
    a.click()

}