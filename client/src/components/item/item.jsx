import numeral from "numeral";
import { Link } from "react-router-dom";
import "./item.css";

const Item = ({ item }) => {
  console.log(item);
  return (
    <div className="item">
      <Link to={`/item/${item.id}`} className="item-link">
        <div className="item-image-container">
          <img src={item.imageURL} alt={item.title} />
        </div>
        <div className="divider"></div>
        <div className="item-details">
          <div className="detail item-price">
            {numeral(item.price).format("$0,0.00")}
          </div>
          <div className="detail item-title">{item.name}</div>
        </div>
      </Link>
    </div>
  );
};

export default Item;
