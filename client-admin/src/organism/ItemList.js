import { useState, useEffect, useRef } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { setItems, createItem, updateItem, deleteItem } from '../store/actionCreator/items';
import { setCategories } from '../store/actionCreator/categories';
import TableRow from './TableRow';
import TableHead from '../molecules/TableHead';
import Form from './Form';
import InputForm from '../molecules/InputForm';
import { toast } from 'react-toastify';

export default function ItemList() {
  const loading = useRef();
  const dispatch = useDispatch();
  const { items, isLoading: itemsLoading, error: itemsError } = useSelector(state => state.items);
  const { categories, isLoading: categoriesLoading, error: categoriesError  } = useSelector(state => state.categories);
  const [formMode, setFormMode] = useState({
    mode: 'creation',
    id: '',
  });
  const initialFormData = {
    name: '',
    description: '',
    price: '',
    imgUrl: '',
    categoryId: '',
  }

  const [formData, setFormData] = useState(initialFormData);
  const [ingredientsFormData, setingredientsFormData] = useState([]);

  useEffect(() => {
    loading.current = toast.info("Loading items...");
    dispatch(setItems())
      .then(() => {
        toast.update(loading.current, {
          type: toast.TYPE.SUCCESS,
          render: 'Loading items completed!',
        });
      })
    dispatch(setCategories());
  }, []);

  useEffect(() => {
    if (itemsError) {
      toast.update(loading.current, {
        type: toast.TYPE.ERROR,
        render: itemsError.message,
      });
    };
  }, [itemsError]);

  const itemUpdate = async (id) => {
    setFormMode({
      mode: 'update',
      id,
    });
    const item = items.filter(item => item.id === id)[0];
    setFormData({
      name: item.name,
      description: item.description,
      price: item.price,
      imgUrl: item.imgUrl,
      categoryId: item.categoryId,
    });
  }

  const handleFormDataChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
       [name]: value,
    });
  }

  const handleIngredientsFormDataChange = (i, e) => {
    const ingredients = [...ingredientsFormData];
    ingredients[i]['name'] = e.target.value;
    setingredientsFormData(ingredients);
  }
  
  const handleAddIngredientField = (e) => {
    e.preventDefault();
    const ingredients = [...ingredientsFormData];
    ingredients.push({ name: '' });
    setingredientsFormData(ingredients);
  };

  const handleRemoveIngredientField = (i, e) => {
    e.preventDefault();
    const ingredients = [...ingredientsFormData];
    ingredients.splice(i, 1);
    setingredientsFormData(ingredients);
  };

  const handleFormSubmit = (e) => {
    e.preventDefault();
    const payload = {
      ...formData,
      ingredients: ingredientsFormData,
    };
    if (formMode.mode === 'creation') {
      loading.current = toast.info("Creating item, please wait...");
      dispatch(createItem(payload))
        .then(() => {
          toast.update(loading.current, {
            type: toast.TYPE.SUCCESS,
            render: 'Item successfully created!',
          });
          setFormData(initialFormData);
          setingredientsFormData([]);
          setFormMode({
            mode: 'creation',
            id: '',
          });
        });
    } else {
      loading.current = toast.info("Updating item, please wait...");
      dispatch(updateItem(formMode.id, formData))
        .then(() => {
          toast.update(loading.current, {
            type: toast.TYPE.SUCCESS,
            render: 'Item successfully updated!',
          });
          setFormData(initialFormData);
          setFormMode({
            mode: 'creation',
            id: '',
          });
        });
    };
  }

  const handleDeleteItem = (id) => {
    loading.current = toast.info("Deleting item, please wait...");
    dispatch(deleteItem(id))
    .then(() => {
      toast.update(loading.current, {
        type: toast.TYPE.SUCCESS,
        render: 'Item successfully deleted!',
      });
    });
  }

  if (itemsLoading || categoriesLoading) {
    return (
      <div>
        <h1>Loading ...</h1>
      </div>
    )
  } else {
    return (
      <div className="flex flex-col max-w-7xl mx-auto">
        <div className="bg-white p-6 mb-3 shadow-md rounded">
          <Form
            buttonText="Submit"
            formData={formData}
            categories={categories}
            onFormDataChange={handleFormDataChange} 
            onFormSubmit={handleFormSubmit}>
              <div>
                {ingredientsFormData.map((ingredient, i) => {
                  return (
                    <>
                      <div className="flex" key={`${ingredient}~${i}`}>
                        <InputForm 
                        key={`${ingredient}~${i}`}
                        name='ingredient'
                        type='text'
                        value={ingredient.name}
                        onFormDataChange={(e) => handleIngredientsFormDataChange(i, e)} 
                        />
                        <button
                          key={`${ingredient}~${i}~${i}`}
                          onClick={(e) => handleRemoveIngredientField(i, e)}
                          className="bg-red-600 text-white active:bg-red-500 text-sm font-bold px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 h-11 mt-auto mb-3 ml-2"
                          style={{ transition: "all .15s ease" }} >
                          Remove
                        </button>
                      </div>
                    </>
                  )
                })}
                <button
                  onClick={(e) => handleAddIngredientField(e)}
                  className="bg-green-600 text-white active:bg-green-500 text-sm font-bold uppercase px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 h-11 mt-auto mb-3"
                  style={{ transition: "all .15s ease" }} >
                  +Add Ingredient
                </button>
              </div>
          </Form>
        </div>
        <div className="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
          <div className="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
            <div className="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <TableHead title="Name" />
                    <TableHead title="Price" />
                    <TableHead title="Category" />
                    <TableHead title="Created By" />
                    <th scope="col" className="relative px-6 py-3">
                      <span className="sr-only">Edit</span>
                    </th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  {items.map((item) => {
                    return <TableRow 
                      key={item.id}
                      item={item} 
                      tableAction={(id) => (
                        <>
                          <button 
                            onClick={() => itemUpdate(id)} 
                            className="text-indigo-600 hover:text-indigo-900 mr-5 font-semibold">
                            Edit
                          </button>
                          <button 
                            onClick={() => handleDeleteItem(id)} 
                            className="text-red-600 hover:text-red-900 font-semibold">
                            Delete
                          </button>
                        </>
                      )}
                    />
                  })}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    )
  }
}
