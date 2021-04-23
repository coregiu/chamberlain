export default {
    name: "menu",
    showUnloginMenu() {
        document.getElementById("unLoginTop").style.display = "block";
        document.getElementById("loginTop").style.display = "none";
    },
    showLoginMenu() {
        document.getElementById("unLoginTop").style.display = "none";
        document.getElementById("loginTop").style.display = "block";
    }
}