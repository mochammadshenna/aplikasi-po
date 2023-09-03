import { useState } from 'react';
import { useNavigate } from "react-router-dom";
import Form from '../organism/Form';
import { toast } from 'react-toastify';

export default function LoginPage() {
  let navigate = useNavigate();
  const [formData, setFormData] = useState({
    email: '',
    password: '',
  })

  const handleFormDataChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
       [name]: value,
    });
  }

  const handleFormSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`${process.env.REACT_APP_BASE_URL}/login`, {
        method: 'POST',
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      })
      if (response.ok) {
        const result = await response.json();
        localStorage.setItem('access_token', result.access_token);
        navigate("/");
        toast.success("Login success!");
      } else {
        const error = await response.json();
        throw {error};
      };
    } catch (error) {
      toast.error(error.message);
      console.log(error);
    }
  }

  return (
    <main>
      <section className="absolute w-full h-full bg-gray-100">
        <div className="container mx-auto px-4 h-full">
          <div className="flex content-center items-center justify-center h-full">
            <div className="w-full lg:w-5/12 md:w-6/12 px-4">
              <div className="relative flex flex-col min-w-0 break-words w-full mb-6 shadow-lg rounded-lg bg-white border-0">
                <div className="flex-auto px-4 lg:px-10 py-10">
                  <div className="text-gray-500 text-center mb-4 font-bold">
                    <small>Sign in with credentials</small>
                  </div>
                  <Form
                    buttonText="Sign in"
                    formData={formData}
                    onFormDataChange={handleFormDataChange} 
                    onFormSubmit={handleFormSubmit} />
                </div>
              </div>
              <div className="flex flex-wrap mt-6 text-center">
                <div className="w-1/2">
                  <a
                    href="#anchor"
                    onClick={e => e.preventDefault()}
                    className="text-gray-700"
                  >
                    <small>Forgot password?</small>
                  </a>
                </div>
                <div className="w-1/2">
                  <a
                    href="#anchor"
                    onClick={e => e.preventDefault()}
                    className="text-gray-700"
                  >
                    <small>Create new account</small>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </main>
  );
}
