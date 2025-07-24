function VerifyWarning() {
  return (
    <div className="flex flex-col justify-center items-center min-h-[50vh] p-4">
      <div className="w-full max-w-md bg-white p-6 rounded-lg shadow-lg border border-gray-200">
        <div className="space-y-3 text-center flex flex-col">
          <h1 className="text-3xl font-bold text-red-600">Error !</h1>  
          <h1 className="text-2xl font-bold text-red-500">Verify Your Email</h1> 
        </div>
      </div>
    </div>
  )
}

export default VerifyWarning
