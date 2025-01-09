export default function ReportPage() {
    return (
      <div className="min-h-screen flex items-center justify-center bg-black text-neon-green">
        <div className="text-center p-8 border-2 border-neon-pink shadow-neon-pink rounded-lg max-w-2xl">
          <h1 className="text-4xl font-bold neon-text">Report a Violation</h1>
          <p className="mt-4 text-lg">Use this form to report any violations.</p>
          <form className="mt-6 p-4 bg-gray-900 border border-neon-green rounded-md text-left">
            <label className="block mb-2 text-neon-green" htmlFor="username">Description of Offender</label>
            <input id="username" type="text" className="w-full p-2 bg-black border border-neon-pink text-neon-green rounded-md" />
            
            <label className="block mt-4 mb-2 text-neon-green" htmlFor="reason">Reason for Report:</label>
            <textarea id="reason" className="w-full p-2 bg-black border border-neon-pink text-neon-green rounded-md"></textarea>
            
            <button type="submit" className="mt-6 px-6 py-3 text-lg font-semibold bg-neon-pink text-black rounded-md shadow-neon-pink transition-transform transform hover:scale-110">
              Submit Report
            </button>
          </form>
        </div>
        
        <style jsx>{`
          .text-neon-green {
            color: #0fffc3;
          }
          .border-neon-pink {
            border-color: #ff00ff;
          }
          .bg-neon-pink {
            background-color: #ff00ff;
          }
          .shadow-neon-pink {
            box-shadow: 0 0 20px #ff00ff;
          }
          .neon-text {
            text-shadow: 0 0 5px #0fffc3, 0 0 10px #0fffc3, 0 0 20px #0fffc3;
          }
        `}</style>
      </div>
    );
  }
  