import { useEffect } from "react";
import { useState } from "react";
import { useParams } from "react-router";
import CustomSpinner from "../components/custom-spinner/custom-spinner";
import ItemList from "../components/item-list/itemList";
import { backendService } from "../services/backend/backend.service";

const ItemListContainer = () => {
  const [products, setProducts] = useState([]);
  const [showSpinner, setShowSpinner] = useState(false);
  const { id: category } = useParams();

  useEffect(() => {
    setShowSpinner(true);
    if (category) {
      backendService
        .getProductsByCategory(category)
        .then((products) => {
          setProducts(products);
        })
        .catch((err) => {
          console.log(`unable to fetch products by category=${category}`, err);
        })
        .finally(() => {
          setShowSpinner(false);
        });
      return;
    }

    backendService
      .getProducts()
      .then((products) => {
        setProducts(products);
      })
      .catch((err) => {
        console.log(`unable to fetch products`, err);
      })
      .finally(() => {
        setShowSpinner(false);
      });
  }, [category]);

  return (
    <div>
      {showSpinner && <CustomSpinner />}
      <ItemList items={products ?? []} />
    </div>
  );
};

export default ItemListContainer;
