export enum Filters {
  vacations = "Vacations",
  courseRelated = "Course Related",
  dn = "DN",
}

export function getTermsFromFilter(filter: Filters): string[] {
  switch (filter) {
    case Filters.vacations:
      return [
        "vacation",
        "vacations",
        "break",
        "breaks",
        "holiday",
        "holidays",
        "weekend",
        "weekends",
      ];
    case Filters.courseRelated:
      return ["course", "courses"];
    case Filters.dn:
      return ["dn"];
  }
}
