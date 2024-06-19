import Image from "next/image";

export type EventImageProps = {
    src: string;
    alt: string;
};

export function EventImage(props: EventImageProps) {
  return (
    <Image
      src={props.src}
      alt={props.alt}
      width={277}
      height={277}
      priority
      className="rounded-2xl"
    />
  );
}
