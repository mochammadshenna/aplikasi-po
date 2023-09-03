export default function SelectForm({label, name, value, options, onFormDataChange}) {
  return (
    <div className="relative w-full mb-3">
    <label
      className="block uppercase text-gray-700 text-xs font-bold mb-2"
      htmlFor="grid-password" >
      {label}
    </label>
    <select 
      value={value || 1}
      onChange={(e) => onFormDataChange(e)}
      name={name}
      className="border-0 px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:ring w-full"
      style={{ transition: "all .15s ease" }}>
      {options.map((option) => {
        return (
          <option value={option.id} key={option.id}>{option.name}</option>
        )
      })}
    </select>
  </div>
  )
}
