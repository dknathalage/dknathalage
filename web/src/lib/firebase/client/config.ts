// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
import { getFirestore } from "firebase/firestore";
import { getAuth } from "firebase/auth";
import { getStorage } from "firebase/storage";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
    apiKey: "AIzaSyDT49PJlstc2PZNxye0EDzMJnvm1mCLfXw",
    authDomain: "dknathalage-89aab.firebaseapp.com",
    projectId: "dknathalage-89aab",
    storageBucket: "dknathalage-89aab.appspot.com",
    messagingSenderId: "496542410425",
    appId: "1:496542410425:web:995b8e77d7e3dd9ef35d68",
    measurementId: "G-VPWVYFDWQG"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
export const analytics = getAnalytics(app);
export const firestore = getFirestore(app);
export const storage = getStorage(app);
export const auth = getAuth(app);