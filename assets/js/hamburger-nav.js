let hamburgerBoolean = false;

function openHamburger() {
  let hamburgerMenu = document.getElementById("menu-hamburger");

  if (!hamburgerBoolean) {
    hamburgerMenu.style.display = "flex";
    hamburgerBoolean = true;
  } else {
    hamburgerMenu.style.display = "none";
    hamburgerBoolean = false;
  }
}
