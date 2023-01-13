import type { Calendar } from "./Calendar";

export interface CalendarByNameResponse {
  currentDay: number;
  totalDays: number;
  currentWeek: number;
  totalWeeks: number;
  calendar: Calendar;
}
