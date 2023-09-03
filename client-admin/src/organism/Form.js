import InputForm from '../molecules/InputForm';
import SelectForm from '../molecules/SelectForm';

export default function Form({buttonText, formData, categories, children, onFormDataChange, onFormSubmit}) {
  let categorySelection = undefined;
  if (categories) {
    categorySelection = <SelectForm
      label='Category'
      name='categoryId'
      value={formData.categoryId}
      options={categories} 
      onFormDataChange={(e) => onFormDataChange(e)}
      />
  }
  return (
    <form onSubmit={(e) => onFormSubmit(e)}>
      {Object.keys(formData).map((key) => {
        const type = key === 'email' || 'password' ? key : 'text';
        if (key.replace(/([a-z0-9])([A-Z])/g, '$1 $2').split(" ")[1] !== 'Id') {
          return <InputForm 
            key={key}
            name={key}
            type={type}
            value={formData[key]}
            onFormDataChange={(e) => onFormDataChange(e)}
          />
        }
      })}
      {categorySelection}
      {children}
      <div className="text-center mt-6">
        <button
          className="bg-gray-900 text-white active:bg-gray-700 text-sm font-bold uppercase px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 w-full"
          type="submit"
          style={{ transition: "all .15s ease" }} >
          {buttonText}
        </button>
      </div>
    </form>
  );
}
