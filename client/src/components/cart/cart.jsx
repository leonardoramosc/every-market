import numeral from "numeral";
import { useContext } from "react";
import { Link, useNavigate } from "react-router-dom";
import { cartContext } from "../../cartContext";
import CartItem from "../cartItem/cartItem";
import "./cart.css";
import { useState } from "react";
import OrderForm from "../order-form/order-form";
import CustomSpinner from "../custom-spinner/custom-spinner";
import OrderInfoModal from "../orderInfoModal/orderInfoModal";
import { backendService } from "../../services/backend/backend.service";

const Cart = () => {
  const { cartItems, clearCartItems } = useContext(cartContext);
  const [showOrderForm, setShowOrderForm] = useState(false);
  const [showSpinner, setShowSpinner] = useState(false);
  const [orderId, setOrderId] = useState(''); 
  const [showOrderId, setShowOrderId] = useState(false);

  const navigate = useNavigate();

  const handleOrderSubmit = async (buyer) => {
    try {
      setShowSpinner(true);
      const newOrderDoc = await backendService.createOrder(buyer, cartItems)
      setOrderId(newOrderDoc.id);
      clearCartItems();
    } catch(err) {
      console.log(err);
    } finally {
      setShowSpinner(false);
      setShowOrderForm(false);
      setShowOrderId(true);
    }
  }

  const renderItems = () => (
    <div className="cart-wrapper">
      <div className="cart-header">
        <h2>Carrito de compras</h2>
        <button onClick={() => setShowOrderForm(true)}>Completar Orden</button>
      </div>
      {cartItems.items.map((item) => (
        <CartItem cartItem={item} key={item.item.id} />
      ))}
      <div className="cart-footer">
        <span>Total</span>
        <span style={{ minWidth: "100px" }}>
          {numeral(cartItems.total).format("$ 0,0.00")}
        </span>
      </div>
    </div>
  );

  const renderEmptyCartMessage = () => (
    <>
      <Link className="back-btn" to="/">Ver Productos</Link>
      <div className="message-container">
        <h2>Tu carrito de compras está vacío</h2>
      </div>
    </>
  );

  return (
    <div className="cart-container">
      {cartItems.items.length > 0 ? renderItems() : renderEmptyCartMessage()}
      <OrderForm show={showOrderForm} onHide={() => setShowOrderForm(false)} onSubmit={handleOrderSubmit} />
      { showSpinner && <CustomSpinner />}
      <OrderInfoModal 
        show={showOrderId} 
        onHide={() => {
          setShowOrderId(false);
          navigate('../');
        }} 
        orderId={orderId} 
      />
    </div>
  );
};

export default Cart;
