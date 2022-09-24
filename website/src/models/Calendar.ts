import type { CalendarEvent } from "./CalendarEvent";

export interface Calendar {
  title: string;
  events: CalendarEvent[];
}
