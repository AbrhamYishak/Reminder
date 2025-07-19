function Setting() {
  return (
    <div className="flex flex-col justify-center items-center min-h-[50vh] p-4">
      <div className="w-full max-w-md bg-white p-6 rounded-lg shadow-lg border border-gray-200">
        <div className="space-y-3 text-center flex flex-col gap-3">
              <label
                className=" self-start text-lg font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                htmlFor="timezone"
              >
                Time Zone
              </label>
              <div className="max-w-xs text-gray-500 flex flex-col gap-3">
                <div className="inset-y-0 my-auto flex items-center">
                  <select className="text-md outline-none rounded-md py-2 px-3 h-full" >
                    <option value="UTC-12:00"> (BIT) UTC-12:00</option>
                    <option value="UTC-11:00"> (NST) UTC-11:00</option>
                    <option value="UTC-10:00"> (HST) UTC-10:00</option>
                    <option value="UTC-09:00"> (AKST) UTC-09:00</option>
                    <option value="UTC-08:00"> (PST) UTC-08:00</option>
                    <option value="UTC-07:00"> (MST) UTC-07:00</option>
                    <option value="UTC-06:00"> (CST) UTC-06:00</option>
                    <option value="UTC-05:00"> (EST) UTC-05:00</option>
                    <option value="UTC-04:00"> (AST) UTC-04:00</option>
                    <option value="UTC-03:00"> (ART) UTC-03:00</option>
                    <option value="UTC-02:00"> (GST) UTC-02:00</option>
                    <option value="UTC-01:00"> (CVT) UTC-01:00</option>
                    <option value="UTC±00:00"> (GMT) UTC±00:00</option>
                    <option value="UTC+01:00"> (CET) UTC+01:00</option>
                    <option value="UTC+02:00"> (EET) UTC+02:00</option>
                    <option value="UTC+03:00"> (EAT) UTC+03:00</option>
                    <option value="UTC+04:00"> (GST) UTC+04:00</option>
                    <option value="UTC+05:00"> (PKT) UTC+05:00</option>
                    <option value="UTC+06:00"> (BST) UTC+06:00</option>
                    <option value="UTC+07:00"> (ICT) UTC+07:00</option>
                    <option value="UTC+08:00"> (CST) UTC+08:00</option>
                    <option value="UTC+09:00"> (JST) UTC+09:00</option>
                    <option value="UTC+10:00"> (AEST) UTC+10:00</option>
                    <option value="UTC+11:00"> (SBT) UTC+11:00</option>
                    <option value="UTC+12:00"> (NZST) UTC+12:00</option>
                  </select>
            </div>

            <div className="flex items-center space-x-2">
              <hr className="flex-grow border-zinc-200 dark:border-zinc-700" />
              <hr className="flex-grow border-zinc-200 dark:border-zinc-700" />
            </div>

            <button
              type="button"
              className="inline-flex  self-center items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2 w-full bg-[#4285F4] text-white"
	          onClick={()=>localStorage.clear}
            >
              <div className="flex items-center justify-center">Log Out</div>
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Setting;
