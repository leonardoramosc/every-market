import { addDoc, collection, doc, getDoc, getDocs, getFirestore, query, Timestamp, updateDoc, where } from "@firebase/firestore";


import { db, firebaseConfig } from "../../firebase/firebase";
import { initializeApp } from "firebase/app";

class FirebaseBackendService {
  constructor() {
    // Initialize Firebase
    const firebaseApp = initializeApp(firebaseConfig);
    const db = getFirestore(firebaseApp);
    this.db = db
  }

  async getProductsByCategory(category) {
    const q = query(
      collection(this.db, "products"),
      where("category", "==", category)
    );
    const querySnapshot = await getDocs(q)

    return querySnapshot.docs.map((doc) => {
      return { ...doc.data(), id: doc.id };
    });
  }

  async getProducts() {
    const productsCollection = collection(db, "products");
    const productDocs = await getDocs(productsCollection)

    return productDocs.docs.map((doc) => {
      return { ...doc.data(), id: doc.id };
    });
  }

  async getProductById(productId) {
    const docRef = doc(db, "products", productId);
    const docSnap = await getDoc(docRef)
    if (!docSnap.exists()) {
      return
    }
    return { id: docSnap.id, ...docSnap.data() }
  }

  async createOrder(cartItems, buyer) {
    const order = {};

    order.date = Timestamp.fromDate(new Date());
    order.buyer = buyer;
    order.total = cartItems.total;
    order.items = cartItems.items.map((cartItem) => {
      const { id, title, price } = cartItem.item;
      return { id, title, price, quantity: cartItem.quantity };
    });

    // crear orden
    const newOrderDoc = await addDoc(collection(db, "orders"), order);

    cartItems.items.forEach((cartItem) => {
      const { item } = cartItem;
      const docRef = doc(db, "products", item.id);
      updateDoc(docRef, { stock: item.stock - cartItem.quantity, sold: item.sold + cartItem.quantity });
    });

    return newOrderDoc
  }
}

// export const firebaseBackendService = new FirebaseBackendService()