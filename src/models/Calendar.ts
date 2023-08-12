import type { CalendarEvent } from "./CalendarEvent";
import { parse, differenceInCalendarDays } from "date-fns";

export interface Calendar {
  title: string;
  events: CalendarEvent[];
}

const getSemesterStartDate = (calendar: Calendar): Date => {
  const sortedEventsByStartDate = [...calendar.events].sort(
    (a, b) =>
      parse(a.start_date, "MMM d, yyyy", new Date()).getTime() -
      parse(b.start_date, "MMM d, yyyy", new Date()).getTime()
  );
  return parse(
    sortedEventsByStartDate[0].start_date,
    "MMM d, yyyy",
    new Date()
  );
};

const getSemesterEndDate = (calendar: Calendar): Date => {
  const sortedEventsByEndDate = [...calendar.events]
    .filter((event) => event.end_date !== null)
    .sort(
      (a, b) =>
        parse(a.end_date!, "MMM d, yyyy", new Date()).getTime() -
        parse(b.end_date!, "MMM d, yyyy", new Date()).getTime()
    );
  return parse(
    sortedEventsByEndDate[sortedEventsByEndDate.length - 1].end_date!,
    "MMM d, yyyy",
    new Date()
  );
};

const currentDayFromSemester = (calendar: Calendar): number => {
  const today = new Date();
  return differenceInCalendarDays(today, getSemesterStartDate(calendar)) + 1;
};

const totalDaysOfSemester = (calendar: Calendar): number => {
  return (
    differenceInCalendarDays(
      getSemesterEndDate(calendar),
      getSemesterStartDate(calendar)
    ) + 1
  );
};

const currentWeekOfSemester = (calendar: Calendar): number => {
  const today = new Date();
  const startSemester = getSemesterStartDate(calendar);
  const daysPassed = differenceInCalendarDays(today, startSemester);
  return Math.floor(daysPassed / 7) + 1;
};

const totalWeeksOfSemester = (calendar: Calendar): number => {
  const startSemester = getSemesterStartDate(calendar);
  const endSemester = getSemesterEndDate(calendar);
  const daysInSemester = differenceInCalendarDays(endSemester, startSemester);
  return Math.floor(daysInSemester / 7) + 1;
};

const isCurrentDayInSemester = (calendar: Calendar): boolean => {
  const today = new Date();
  // Setting the hours, minutes, seconds, and milliseconds to 0 for accurate comparison
  today.setHours(0, 0, 0, 0);

  const firstDate = getSemesterStartDate(calendar);
  const lastDate = getSemesterEndDate(calendar);

  return today >= firstDate && today <= lastDate;
};

export {
  currentDayFromSemester,
  totalDaysOfSemester,
  currentWeekOfSemester,
  totalWeeksOfSemester,
  isCurrentDayInSemester,
};
