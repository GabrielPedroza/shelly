export default function Landing() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-black text-neon-green">
      <div className="text-center p-8 border-2 border-neon-pink shadow-neon-pink rounded-lg max-w-2xl">
        <h1 className="text-5xl font-bold neon-text">Welcome to Shellhacks</h1>
        <p className="mt-4 text-lg">Experience the cyberpunk revolution. The future is now.</p>
        <button className="mt-6 px-6 py-3 text-lg font-semibold bg-neon-pink text-black rounded-md shadow-neon-pink transition-transform transform hover:scale-110">
          Register
        </button>
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
