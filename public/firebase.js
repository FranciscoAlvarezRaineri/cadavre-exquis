import { initializeApp } from "https://www.gstatic.com/firebasejs/10.1.0/firebase-app.js";

import {
  getAuth,
  signInWithEmailAndPassword,
  signOut
} from "https://www.gstatic.com/firebasejs/10.1.0/firebase-auth.js";

const firebaseConfig = {
  apiKey: "AIzaSyCYMDy_Rf1o82MmWVyI0p0fRwyJ2w2dcTs",
  authDomain: "cadavre-exquis-9c7af.firebaseapp.com",
  projectId: "cadavre-exquis-9c7af",
  storageBucket: "cadavre-exquis-9c7af.appspot.com",
  messagingSenderId: "950495760420",
  appId: "1:950495760420:web:57784044dd9be8f95548b1",
  measurementId: "G-80DH0CRH6R"
}

const app = initializeApp(firebaseConfig);

const auth = getAuth(app);

function signIn(email, password) {
  signInWithEmailAndPassword(auth, email, password)
    .then((userCredential) => {
      const user = userCredential.user;
      document.cookie = `accessToken=${user.accessToken}`;

      htmx.ajax("GET", '/home', '#main')
    })
    .catch((error) => {
      document.getElementById("msg").innerText = "invalid credentials, please try again"
      document.getElementById("email").value = ""
      document.getElementById("password").value = ""
    });
}

function signOff() {
  signOut(auth)
    .then((userCredential) => {
      document.cookie = `accessToken=`;
      htmx.ajax("GET", '/home', '#main')
    })
    .catch((error) => {
      htmx.ajax("GET", '/home', '#main')
    });
}

window.signIn = signIn
window.signOff = signOff