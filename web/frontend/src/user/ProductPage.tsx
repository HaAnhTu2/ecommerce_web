import React, {useEffect, useState} from "react";
import { getProducts, deleteProduct } from "../services/api";
import { Product } from "../types/product";

const ProductPage:React.FC=()=>{
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
    const handleDeleteProduct =async (id: string)=>{
        try{
            await deleteProduct(id)
            setProducts(products.filter(product=>product._id!=id))
        }catch(error){
            console.error('error deleting product:',error);
        }
    }
    return(
        <div>
            <div>
                <h4>Product</h4>
            </div>
            <table>
                <thead>
                <tr>
                    <th>Product Name</th>
                    <th>Description</th>
                    <th>Price</th>
                    <th>Stock</th>
                    <th>Category</th>
                    <th>Action</th>
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
                            <td>
                                <button type="submit" onClick={()=> handleDeleteProduct(product._id)}>Delete</button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}
export default ProductPage;
