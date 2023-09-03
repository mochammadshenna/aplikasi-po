export default function InputForm({name, value, type, onFormDataChange}) {
  const label = name.replace(/([a-z0-9])([A-Z])/g, '$1 $2').replace(/^./, (str) => { return str.toUpperCase(); });
  return (
    <div className="relative w-full mb-3">
      <label
        className="block uppercase text-gray-700 text-xs font-bold mb-2"
        htmlFor="grid-password" >
        {label}
      </label>
      <input
        value={value}
        onChange={(e) => onFormDataChange(e)}
        name={name}
        type={type}
        className="border-0 px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:ring w-full"
        placeholder={label}
        style={{ transition: "all .15s ease" }} />
    </div>
  );
};
