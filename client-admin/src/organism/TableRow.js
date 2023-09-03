export default function TableRow({item, tableAction}) {
  return (
    <tr>
      <td className="px-6 py-4 whitespace-nowrap">
        <div className="flex items-center">
          <div className="flex-shrink-0 h-10 w-10">
            <img className="h-10 w-10 rounded-full" src={item.imgUrl} alt={item.name}/>
          </div>
          <div className="ml-4">
            <div className="text-sm font-medium text-gray-900">{item.name}</div>
            {/* <div className="text-sm text-gray-500">{item.description}</div> */}
          </div>
        </div>
      </td>
      <td className="px-6 py-4 whitespace-nowrap">
        <div className="text-sm text-gray-900">Rp. {item.price}</div>
      </td>
      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{item.Category?.name}</td>
      <td className="px-6 py-4 whitespace-nowrap">
        <div className="text-sm text-gray-900">{item.User?.username}</div>
        <div className="text-sm text-gray-500">{item.User?.email}</div>
      </td>
      <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
        {tableAction(item.id)}
      </td>
    </tr>
  )
}
