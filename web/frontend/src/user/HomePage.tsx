import React, {useEffect, useState} from "react";
import { getProducts } from "../services/api";
import { Product } from "../types/product";

const UserHome:React.FC=()=>{
    const [products,setProducts] =useState<Product[]>([]);
    const [loading, setloading] = useState(true);

    useEffect(()=>{
        const fetchProducts= async()=>{
            try{
                const fetchedProducts =await getProducts();
                setProducts(fetchedProducts);
                setloading(false);
            }catch(error){
                console.log('error fetching products:',error);
                setloading(false);
            }
        };
        fetchProducts();
    },[]);
    if(loading){
        return <div>Loading...</div>;
    }

    return(
        <div className="container">
            <table>
                <thead>
                <tr>
                    <th>Product Name</th>
                    <th>Description</th>
                    <th>Price</th>
                    <th>Stock</th>
                    <th>Category</th>
                </tr>
                </thead>
                <tbody>
                    {products?.map(product=>(
                        <tr key={product._id}>
                            <td>{product.name}</td>
                            <td>{product.description}</td>
                            <td>{product.price}</td>
                            <td>{product.stock}</td>
                            <td>{product.category}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}
export default UserHome;

