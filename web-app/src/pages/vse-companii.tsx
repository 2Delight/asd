import {NextPage} from "next";
import { useRouter } from "next/router";

export default function About() {
    const { asPath, pathname } = useRouter()
    console.log(asPath, pathname)
    return <div>About</div>
}