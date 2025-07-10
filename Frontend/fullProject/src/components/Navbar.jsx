import React, { useRef, useState ,useContext} from "react";
import { motion } from "framer-motion";
function Nav() {
  const [position, setPosition] = useState({
    left: 0,
    width: 0,
    opacity: 0,
  });
  return (
    <ul
      className="margin-auto flex w-fit rounded-full border-white bg-black/30  text-black z-10"
      onMouseLeave={() => setPosition((pv) => ({ ...pv, opacity: 0 }))}
    >
      <Tab setPosition={setPosition}> 
  Home
</Tab>
      <Tab setPosition={setPosition}>
  Active Messages
</Tab>
      <Tab setPosition={setPosition}>
  Sent Messages
</Tab>
<Cursor position={position} />
    </ul>
  );
}

const Tab = ({ children, setPosition }) => {
  const ref = useRef(null);

  return (
    <li
      ref={ref}
      onMouseEnter={() => {
        if (!ref.current) return;

        const { width } = ref.current.getBoundingClientRect();
        setPosition({
          width,
          opacity: 1,
          left: ref.current.offsetLeft,
        });
      }}
      className="relative z-10 block cursor-pointer px-3 py-1.5 text-xs uppercase text-white mix-blend-difference md:px-5 md:py-3 md:text-base"
    >
      {children}
    </li>
  );
};

const Cursor = ({ position }) => {
  return (
    <motion.li
      animate={position}
      className="absolute z-0 h-7 rounded-full bg-black md:h-12"
    />
  );
};

export default Nav;

