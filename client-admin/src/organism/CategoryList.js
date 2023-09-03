import { useState, useEffect, useRef } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { setCategories, createCategory, deleteCategory } from '../store/actionCreator/categories';
import TableRowCategory from './TableRowCategory';
import TableHead from '../molecules/TableHead';
import Form from './Form';
import { toast } from 'react-toastify';

export default function CategoryList() {
  const loading = useRef();
  const dispatch = useDispatch();
  const { categories, isLoading, error } = useSelector(state => state.categories);
  const initialFormData = {
    name: '',
  }

  const [formData, setFormData] = useState(initialFormData);

  useEffect(() => {
    dispatch(setCategories());
  }, [])

  const handleFormDataChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  }

  const handleFormSubmit = (e) => {
    e.preventDefault();
    loading.current = toast.info("Creating category, please wait...");
    dispatch(createCategory(formData))
      .then(() => {
        toast.update(loading.current, {
          type: toast.TYPE.SUCCESS,
          render: 'Category successfully created!',
        });
      });
    setFormData(initialFormData);
  }

  const handleDeleteCategory = (id) => {
    loading.current = toast.info("Deleting category, please wait...");
    dispatch(deleteCategory(id))
      .then(() => {
        toast.update(loading.current, {
          type: toast.TYPE.SUCCESS,
          render: 'Category successfully deleted!',
        });
      });
  }

  if (isLoading) {
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
            onFormDataChange={handleFormDataChange} 
            onFormSubmit={handleFormSubmit} />
        </div>
        <div className="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
          <div className="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
            <div className="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <TableHead title="Name" />
                    <TableHead title="Created At" />
                    <TableHead title="Updated At" />
                    <th scope="col" className="relative px-6 py-3">
                      <span className="sr-only">Delete</span>
                    </th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  {categories.map((category) => {
                    return <TableRowCategory
                      key={category.id}
                      category={category} 
                      tableAction={(id) => (
                        <button 
                          onClick={() => handleDeleteCategory(id)} 
                          className="text-red-600 hover:text-red-900 font-semibold">
                          Delete
                        </button>
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
